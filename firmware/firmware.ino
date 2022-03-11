#include <ESP8266WiFi.h>
#include <ESP8266mDNS.h>
#include <ESPAsyncTCP.h>
#include <ESPAsyncWebServer.h>
#include <ESPAsyncWiFiManager.h>
#include "form.h"

#define HOSTNAME "squirtianna"
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

  // WiFi.mode(WIFI_STA);
  // WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  // if (WiFi.waitForConnectResult() != WL_CONNECTED) {
  //     Serial.printf("WiFi Failed!\n");
  //     return;
  // }
  
  Serial.print("Wifi ready: ");
  Serial.println(WiFi.localIP());
}

void setupWebserver(){

  server.on("/", HTTP_GET, [] (AsyncWebServerRequest *request) {
    request->send(200, "text/html", indexhtml);
  });

  server.on("/squirt", HTTP_GET, [] (AsyncWebServerRequest *request) {
    if (request->hasParam("duration")) {
      int duration = request->getParam("duration")->value().toInt();
      del = duration;
    }
    request->redirect("/");
  });

  server.onNotFound([](AsyncWebServerRequest *request){
    request->redirect("https://github.com/tlanfer/squirtianna");
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
  digitalWrite(12, HIGH);
  digitalWrite(13, HIGH);
  digitalWrite(14, HIGH);
  digitalWrite(15, HIGH);
  delay(t);
  digitalWrite(12, LOW);
  digitalWrite(13, LOW);
  digitalWrite(14, LOW);
  digitalWrite(15, LOW);
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
