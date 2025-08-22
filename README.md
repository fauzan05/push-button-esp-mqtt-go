# ESP8266 MQTT Button Project

A simple IoT project that sends button press events from ESP8266 to a Golang backend via MQTT protocol.

## 🚀 Features

- Real-time button press detection on ESP8266
- MQTT messaging between ESP8266 and Golang
- JSON formatted data transmission
- Debounced button input handling
- Local MQTT broker communication
- Cross-platform Golang receiver

## 🛠️ Hardware Requirements

- ESP8266 Development Board (Lolin/Wemos D1 Mini)
- Push Button Switch
- Jumper wires
- Breadboard (optional)

## 🔧 Software Requirements

- [PlatformIO](https://platformio.org/) for ESP8266 development
- [Golang](https://golang.org/) (v1.19 or higher)
- [Mosquitto MQTT Broker](https://mosquitto.org/)
- macOS/Linux/Windows

## 📋 Project Structure

```
esp8266-mqtt-button/
├── src/
│   └── main.cpp          # ESP8266 source code
├── go/
│   ├── go.mod            # Go module dependencies
│   └── main.go           # Golang MQTT receiver
├── include/
├── lib/
├── test/
├── platformio.ini        # PlatformIO configuration
└── README.md
```

## ⚡ Hardware Setup

### Wiring Diagram
```
ESP8266 (Lolin)    Push Button
D1 (GPIO5)     →   Pin 1
GND            →   Pin 2
```

**Note:** ESP8266 has internal pull-up resistor, so no external resistor needed.

## 🔨 Software Setup

### 1. Clone Repository
```bash
git clone https://github.com/yourusername/esp8266-mqtt-button.git
cd esp8266-mqtt-button
```

### 2. Setup MQTT Broker (macOS)
```bash
# Install Mosquitto
brew install mosquitto

# Start MQTT broker
mosquitto -v

# Or run as service
brew services start mosquitto
```

### 3. ESP8266 Setup
```bash
# Initialize PlatformIO project (if not cloned)
pio project init --board d1_mini

# Update WiFi credentials in src/main.cpp
# Change YOUR_WIFI_SSID, YOUR_WIFI_PASSWORD, YOUR_MAC_IP_ADDRESS

# Build and upload
pio run --target upload

# Monitor serial output
pio device monitor
```

### 4. Golang Backend Setup
```bash
# Navigate to golang directory
cd golang

# Install dependencies
go mod tidy

# Run the receiver
go run main.go
```

## 📊 Data Format

ESP8266 sends JSON data via MQTT:

```json
{
  "device": "esp8266_lolin",
  "action": "button_pressed", 
  "timestamp": 12345,
  "status": "active"
}
```

## 🎯 Usage

1. **Start MQTT Broker**
   ```bash
   mosquitto -v
   ```

2. **Run Golang Receiver**
   ```bash
   cd golang
   go run main.go
   ```

3. **Upload ESP8266 Code**
   ```bash
   pio run --target upload
   ```

4. **Press the button** and watch real-time data in terminal!

## 📈 Expected Output

### Golang Console:
```
🚀 Starting Golang MQTT Button Receiver...
✅ Connected to MQTT broker!
📡 Subscribed to topic: esp8266/button
⏳ Waiting for button presses... (Press Ctrl+C to exit)

=== NEW MESSAGE RECEIVED ===
Topic: esp8266/button
Device: esp8266_lolin
Action: button_pressed
ESP8266 Timestamp: 12345 ms
Status: active
Received at: 2025-08-22 14:30:25
============================
🔘 Processing button press from esp8266_lolin
🎉 Button was pressed! Executing custom logic...
```

### ESP8266 Serial Monitor:
```
WiFi connected
IP address: 192.168.1.100
connected
Button pressed!
Message sent successfully!
```

## 🐛 Troubleshooting

### ESP8266 Issues
| Problem | Solution |
|---------|----------|
| WiFi not connecting | Check SSID/password, signal strength |
| MQTT connection failed | Verify Mac IP address, check broker status |
| Button not responding | Check wiring, enable pull-up resistor |

### Golang Issues
| Problem | Solution |
|---------|----------|
| No messages received | Check topic name, restart broker |
| Connection refused | Ensure Mosquitto is running |
| Import errors | Run `go mod tidy` |

### Debug Commands
```bash
# Monitor all MQTT topics
mosquitto_sub -h localhost -t "#" -v

# Test MQTT manually
mosquitto_pub -h localhost -t esp8266/button -m '{"test":"message"}'

# Check network connectivity
ping YOUR_ESP8266_IP
```

## 🔧 Configuration

### ESP8266 Configuration
Update these variables in `src/main.cpp`:
```cpp
const char* ssid = "YOUR_WIFI_SSID";
const char* password = "YOUR_WIFI_PASSWORD"; 
const char* mqtt_server = "YOUR_MAC_IP_ADDRESS";
const char* mqtt_topic = "esp8266/button";
```

### Golang Configuration
Update these constants in `golang/main.go`:
```go
const (
    broker   = "localhost"
    port     = 1883
    topic    = "esp8266/button"
    clientID = "golang-mqtt-client"
)
```

## 🚀 Future Enhancements

- [ ] Web dashboard for monitoring
- [ ] Database logging (SQLite/PostgreSQL)
- [ ] Multiple sensor support
- [ ] Push notifications
- [ ] RESTful API
- [ ] Docker containerization
- [ ] Two-way communication (control ESP8266 from Golang)

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Your Name**
- GitHub: [@yourusername](https://github.com/yourusername)
- Email: your.email@example.com

## 🙏 Acknowledgments

- [PlatformIO](https://platformio.org/) for ESP8266 development environment
- [Eclipse Paho](https://github.com/eclipse/paho.mqtt.golang) for Golang MQTT client
- [Mosquitto](https://mosquitto.org/) for MQTT broker
- ESP8266 Community for excellent documentation

---

⭐ If this project helped you, please consider giving it a star!