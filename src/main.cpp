#include <ESP8266WiFi.h>
#include <PubSubClient.h>

// WiFi credentials
const char* ssid = "YOUR_WIFI_SSID";
const char* password = "YOUR_WIFI_PASSWORD";

// MQTT Broker settings
const char* mqtt_server = "YOUR_MAC_IP_ADDRESS"; // Ganti dengan IP Mac Anda
const int mqtt_port = 1883;
const char* mqtt_topic = "esp8266/button";

// Button settings
const int buttonPin = D1;  // GPIO5 pada Lolin
bool lastButtonState = HIGH;
bool buttonState = HIGH;
unsigned long lastDebounceTime = 0;
unsigned long debounceDelay = 50;

WiFiClient espClient;
PubSubClient client(espClient);

void setup() {
  Serial.begin(115200);
  
  // Setup button
  pinMode(buttonPin, INPUT_PULLUP);
  
  // Connect to WiFi
  setup_wifi();
  
  // Setup MQTT
  client.setServer(mqtt_server, mqtt_port);
  
  Serial.println("ESP8266 Button to MQTT Ready!");
}

void setup_wifi() {
  delay(10);
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void reconnect() {
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    
    // Create a random client ID
    String clientId = "ESP8266Client-";
    clientId += String(random(0xffff), HEX);
    
    if (client.connect(clientId.c_str())) {
      Serial.println("connected");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
    }
  }
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  // Read button with debouncing
  int reading = digitalRead(buttonPin);
  
  if (reading != lastButtonState) {
    lastDebounceTime = millis();
  }
  
  if ((millis() - lastDebounceTime) > debounceDelay) {
    if (reading != buttonState) {
      buttonState = reading;
      
      // Button pressed (LOW because of INPUT_PULLUP)
      if (buttonState == LOW) {
        Serial.println("Button pressed!");
        
        // Create JSON message
        String message = "{";
        message += "\"device\":\"esp8266_lolin\",";
        message += "\"action\":\"button_pressed\",";
        message += "\"timestamp\":" + String(millis()) + ",";
        message += "\"status\":\"active\"";
        message += "}";
        
        // Publish to MQTT
        if (client.publish(mqtt_topic, message.c_str())) {
          Serial.println("Message sent successfully!");
          Serial.println("Message: " + message);
        } else {
          Serial.println("Failed to send message");
        }
        
        // LED indication (if available)
        digitalWrite(LED_BUILTIN, LOW);
        delay(200);
        digitalWrite(LED_BUILTIN, HIGH);
      }
    }
  }
  
  lastButtonState = reading;
  delay(10);
}