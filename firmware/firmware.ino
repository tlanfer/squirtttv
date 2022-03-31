#include <FS.h>
#include <ESP8266WiFi.h>
#include <ESP8266mDNS.h>
#include <ESPAsyncTCP.h>
#include <ESPAsyncWebServer.h>
#include <ESPAsyncWiFiManager.h>
#include <SPIFFSEditor.h>
#include <AsyncElegantOTA.h>
#include <Hash.h>
#include <elegantWebpage.h>

#define HOSTNAME "squirtttv"

#define ADMIN_USERNAME "admin"
#define ADMIN_PASSWORD "admin"

#define PUMP_PIN 12

AsyncWebServer server(80);
DNSServer dns;

int del = 0;

void setupOutput(){
  pinMode(PUMP_PIN, OUTPUT);
  digitalWrite(PUMP_PIN, LOW);
}

void setupSerial(){
  Serial.begin(115200);
  Serial.println("Serial ready");
}

void setupWifi(){
  AsyncWiFiManager wifiManager(&server,&dns);
  wifiManager.autoConnect(HOSTNAME);
  
  Serial.print("Wifi ready: ");
  Serial.println(WiFi.localIP());
}

void setupWebserver(){
  SPIFFS.begin();
  server.serveStatic("/", SPIFFS, "/").setDefaultFile("index.html");
  server.addHandler(new SPIFFSEditor(ADMIN_USERNAME,ADMIN_PASSWORD));

  server.on("/squirt", HTTP_GET, [] (AsyncWebServerRequest *request) {
    if (request->hasParam("duration")) {
      int duration = request->getParam("duration")->value().toInt();
      del = duration;
    }
    request->redirect("/");
  });

  AsyncElegantOTA.begin(&server);    // Start ElegantOTA

  server.onNotFound([](AsyncWebServerRequest *request){
    request->redirect("https://github.com/tlanfer/squirtttv");
  });

  
  server.begin();
  Serial.println("Webserver ready");
}

void setupMdns(){
  if (!MDNS.begin(HOSTNAME)) {
    Serial.println("Error setting up MDNS responder!");
  }
  MDNS.addService(HOSTNAME, "tcp", 80); // Announce esp tcp service on port 8080
  
  Serial.println("mDNS ready");
}

void setup() {
  setupOutput();
  setupSerial();
  setupWifi();
  setupWebserver();
  setupMdns();
}

void squirt(int t) {
  digitalWrite(PUMP_PIN, HIGH);
  delay(t);
  digitalWrite(PUMP_PIN, LOW);
}

void loop() {
  MDNS.update();
  if( del > 0) {
    Serial.print("Squirt for ");
    Serial.print(del);
    Serial.println("ms");
    squirt(del);
    del = 0;
  }
}
