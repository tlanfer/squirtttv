#include <FS.h>
#include <WiFi.h>
#include <LittleFS.h>
#include <AsyncFsWebServer.h>
#include <ESPmDNS.h>

#define HOSTNAME "squirtttv"

#define ADMIN_USERNAME "admin"
#define ADMIN_PASSWORD "admin"

#define PUMP_PIN D4
#define FORMAT_LITTLEFS_IF_FAILED true

AsyncFsWebServer server(80, LittleFS, "myServer");
bool captiveRun = false;

int del = 0;

void setupOutput() {
  pinMode(PUMP_PIN, OUTPUT);
  digitalWrite(PUMP_PIN, LOW);
}

void setupSerial() {
  Serial.begin(115200);
  Serial.println("Serial ready");
}

void setupFS(){
  if (LittleFS.begin()){
    File root = LittleFS.open("/", "r");
    File file = root.openNextFile();
    while (file){
      Serial.printf("FS File: %s, size: %d\n", file.name(), file.size());
      file = root.openNextFile();
    }
  } else {
    Serial.println("ERROR on mounting filesystem. It will be reformatted!");
    LittleFS.format();
    ESP.restart();
  }
  Serial.println("FS ready");
}

void getFsInfo(fsInfo_t* fsInfo) {
	fsInfo->fsName = "LittleFS";
	fsInfo->totalBytes = LittleFS.totalBytes();
	fsInfo->usedBytes = LittleFS.usedBytes();  
}

void setupWebserver() {
  IPAddress myIP = server.startWiFi(15000);
  if (!myIP) {
    Serial.println("\n\nNo WiFi connection, start AP and Captive Portal\n");
    myIP = WiFi.softAPIP();
    Serial.print("My IP 1 ");
    Serial.println(myIP.toString());
    server.startCaptivePortal("ESP_AP", "123456789", "/setup");
    myIP = WiFi.softAPIP();
    Serial.print("\nMy IP 2 ");
    Serial.println(myIP.toString());
  }
  server.setSetupPageTitle("Simple Async FS Captive Web Server");
  server.serveStatic("/", LittleFS, "/").setDefaultFile("index.html");

  server.on("/identify", HTTP_GET, [](AsyncWebServerRequest *request) {
    AsyncWebServerResponse *response = request->beginResponse(200, "text/plain", "This is a streamer squirter");
    response->addHeader("Server", "squirtttv/2.0");

    request->send(response);
  });

  server.on("/squirt", HTTP_GET, [](AsyncWebServerRequest *request) {
    if (request->hasParam("duration")) {
      int duration = request->getParam("duration")->value().toInt();
      del = duration;
    }
    request->redirect("/");
  });

  server.onNotFound([](AsyncWebServerRequest *request) {
    request->redirect("https://github.com/tlanfer/squirtttv");
  });

  server.enableFsCodeEditor();
  server.setFsInfoCallback(getFsInfo);

  server.init();
  Serial.println("Webserver ready");
}

void setupMdns(){
  if (!MDNS.begin("esp32")) {
      Serial.println("Error setting up MDNS responder!");
      while(1) {
          delay(1000);
      }
  }
  
  MDNS.addService("squirtttv", "tcp", 80);
}

void setup() {
  setupOutput();
  setupSerial();
  setupFS();
  setupWebserver();
  setupMdns();
}

void squirt(int t) {
  digitalWrite(PUMP_PIN, HIGH);
  int squirtTo = millis() + t;
  while (millis() < squirtTo){
    delay(1);
  }
  digitalWrite(PUMP_PIN, LOW);
}

void loop() {
  if (captiveRun)
    server.updateDNS();
  
  if (del > 0) {
    Serial.print("Squirt for ");
    Serial.print(del);
    Serial.println("ms");
    squirt(del);
    Serial.println("done squirting");
    del = 0;
  }
}
