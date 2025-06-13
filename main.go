package main

import (
	"embed"
	"fmt"
	"math/rand"
	"net"
	"nws/internal/model"
	"nws/internal/repository"
	"nws/internal/service"
	"nws/internal/storage"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"gorm.io/gorm"
)

//go:embed all:frontend/dist
var assets embed.FS

func simulateLogs(svc *service.Service) {
	devices, err := svc.GetAllDevices()
	if err != nil {
		fmt.Printf("Error getting devices: %v\n", err)
		return
	}

	for {
		// Выбираем случайное устройство
		if len(devices) > 0 {
			randomDevice := devices[rand.Intn(len(devices))]

			// Генерируем лог для устройства
			_, err := svc.SimulateDeviceLog(randomDevice.ID)
			if err != nil {
				fmt.Printf("Error simulating log for device %d: %v\n", randomDevice.ID, err)
			}
		}

		// Ждем случайное время от 5 до 15 секунд перед следующей генерацией
		sleepTime := time.Duration(rand.Intn(10)+5) * time.Second
		time.Sleep(sleepTime)
	}
}

func simulateMetrics(svc *service.Service) {
	devices, err := svc.GetAllDevices()
	if err != nil {
		fmt.Printf("Error getting devices: %v\n", err)
		return
	}

	// Интервалы обновления метрик для разных устройств
	intervals := []string{"1s", "5s", "10s", "30s", "1m", "5m"}
	deviceIntervals := make(map[int]string)

	// Назначаем случайные интервалы устройствам
	for _, device := range devices {
		deviceIntervals[device.ID] = intervals[rand.Intn(len(intervals))]
	}

	// Запускаем горутину для очистки старых метрик каждый час
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			if err := svc.CleanupOldMetrics(24 * time.Hour); err != nil {
				fmt.Printf("Error cleaning up old metrics: %v\n", err)
			}
		}
	}()

	// Основной цикл генерации метрик
	for {
		currentTime := time.Now()

		for _, device := range devices {
			interval := deviceIntervals[device.ID]

			// Проверяем, нужно ли обновлять метрики для этого устройства
			shouldUpdate := false
			switch interval {
			case "1s":
				shouldUpdate = true
			case "5s":
				shouldUpdate = currentTime.Second()%5 == 0
			case "10s":
				shouldUpdate = currentTime.Second()%10 == 0
			case "30s":
				shouldUpdate = currentTime.Second()%30 == 0
			case "1m":
				shouldUpdate = currentTime.Second() == 0
			case "5m":
				shouldUpdate = currentTime.Second() == 0 && currentTime.Minute()%5 == 0
			}

			if shouldUpdate {
				_, err := svc.SimulateDeviceMetric(device.ID, interval)
				if err != nil {
					fmt.Printf("Error simulating metrics for device %d: %v\n", device.ID, err)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func simulateDeviceCreation(svc *service.Service) {
	// Типы устройств и их характеристики
	deviceTypes := []map[string]string{
		{
			"type":   "Router",
			"vendor": "Cisco",
			"model":  "ISR 4321",
		},
		{
			"type":   "Switch",
			"vendor": "Juniper",
			"model":  "EX3400",
		},
		{
			"type":   "Firewall",
			"vendor": "Palo Alto",
			"model":  "PA-820",
		},
		{
			"type":   "LoadBalancer",
			"vendor": "F5",
			"model":  "BIG-IP 2000s",
		},
	}

	// Запускаем горутину для создания новых устройств
	go func() {
		baseIP := net.ParseIP("192.168.1.0")
		ip := make(net.IP, len(baseIP))
		copy(ip, baseIP)
		deviceCounter := 1

		for {
			// Создаем новое устройство каждые 30-60 секунд
			sleepTime := time.Duration(30+rand.Intn(30)) * time.Second
			time.Sleep(sleepTime)

			// Увеличиваем последний октет IP-адреса
			ip[len(ip)-1] = byte(deviceCounter)
			deviceType := deviceTypes[rand.Intn(len(deviceTypes))]

			// Создаем новое устройство
			deviceData := map[string]string{
				"hostname":   fmt.Sprintf("%s-%d.local", strings.ToLower(deviceType["type"]), deviceCounter),
				"ip_address": ip.String(),
				"type":       deviceType["type"],
				"vendor":     deviceType["vendor"],
				"model":      deviceType["model"],
				"serial":     fmt.Sprintf("SN%d", rand.Int63()),
				"os_version": fmt.Sprintf("%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)),
				"status":     "up",
			}

			device, err := svc.CreateDevice(deviceData)
			if err != nil {
				fmt.Printf("Error creating device: %v\n", err)
				continue
			}

			fmt.Printf("Created new device: %s (%s)\n", deviceData["hostname"], deviceData["ip_address"])

			// Генерируем начальные логи для нового устройства
			for i := 0; i < 3; i++ {
				_, err := svc.SimulateDeviceLog(device.ID)
				if err != nil {
					fmt.Printf("Error generating initial logs for device %d: %v\n", device.ID, err)
				}
				time.Sleep(time.Second)
			}

			// Генерируем начальные метрики для нового устройства
			intervals := []string{"1s", "5s", "10s", "30s", "1m", "5m"}
			randomInterval := intervals[rand.Intn(len(intervals))]
			_, err = svc.SimulateDeviceMetric(device.ID, randomInterval)
			if err != nil {
				fmt.Printf("Error generating initial metrics for device %d: %v\n", device.ID, err)
			}

			deviceCounter++
		}
	}()
}

//func simulateConnections(svc *service.Service) {
//	protocols := []string{"TCP", "UDP", "ICMP", "HTTP", "HTTPS", "SSH", "FTP"}
//	ports := map[string][]int{
//		"HTTP":  {80, 8080},
//		"HTTPS": {443, 8443},
//		"SSH":   {22},
//		"FTP":   {21, 20},
//		"TCP":   {135, 139, 445},
//		"UDP":   {53, 67, 68},
//		"ICMP":  {0},
//	}
//
//	go func() {
//		for {
//			// Получаем все соединения
//			connections, err := svc.GetAllConnections()
//			if err != nil {
//				fmt.Printf("Error getting connections: %v\n", err)
//				time.Sleep(5 * time.Second)
//				continue
//			}
//
//			// Обновляем каждое соединение
//			for _, conn := range connections {
//				// Симулируем случайные изменения в метриках соединения
//				conn.Protocol = protocols[rand.Intn(len(protocols))]
//				conn.Port = ports[conn.Protocol][rand.Intn(len(ports[conn.Protocol]))]
//				conn.Status = "active"
//
//				// Симулируем сетевые метрики
//				conn.Latency = rand.Intn(100) + 1 // 1-100ms
//				conn.PacketLoss = rand.Intn(5)    // 0-5%
//
//				// Симулируем трафик
//				basePackets := rand.Intn(1000) + 500 // 500-1500 пакетов
//				conn.PacketsSent = basePackets
//
//				// Рассчитываем потери и ошибки
//				lostPackets := int(float64(basePackets) * float64(conn.PacketLoss) / 100.0)
//				conn.PacketsLost = lostPackets
//				conn.PacketsReceived = basePackets - lostPackets
//
//				// Симулируем различные проблемы с пакетами
//				conn.PacketsReordered = rand.Intn(10)   // 0-10 пакетов
//				conn.PacketsDuplicated = rand.Intn(5)   // 0-5 пакетов
//				conn.PacketsCorrupted = rand.Intn(3)    // 0-3 пакета
//				conn.PacketsReassembled = rand.Intn(15) // 0-15 пакетов
//
//				// Добавляем случайные проблемы
//				if rand.Float64() < 0.05 { // 5% шанс проблем
//					switch rand.Intn(3) {
//					case 0:
//						conn.Status = "degraded"
//						conn.Latency += rand.Intn(200)   // Добавляем 0-200ms к задержке
//						conn.PacketLoss += rand.Intn(10) // Добавляем 0-10% к потерям
//					case 1:
//						conn.Status = "unstable"
//						conn.PacketsReordered += rand.Intn(50)  // Значительное переупорядочивание
//						conn.PacketsDuplicated += rand.Intn(30) // Много дубликатов
//					case 2:
//						conn.Status = "error"
//						conn.PacketsCorrupted += rand.Intn(20) // Много поврежденных пакетов
//						conn.PacketLoss += rand.Intn(15)       // Высокие потери
//					}
//				}
//
//				// Обновляем соединение в базе данных
//				err := svc.UpdateConnection(&conn)
//				if err != nil {
//					fmt.Printf("Error updating connection %d: %v\n", conn.ID, err)
//					continue
//				}
//
//				// Добавляем небольшую задержку между обновлениями
//				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
//			}
//
//			// Ждем перед следующим циклом обновлений
//			time.Sleep(5 * time.Second)
//		}
//	}()
//}

func main() {
	var err error

	db := storage.NewStorage("nws.db")

	models := []interface{}{
		&model.User{},
		&model.NetworkDevice{},
		&model.Connection{},
		&model.NetworkTopology{},
		&model.TopologyElement{},
		&model.NetworkLog{},
		&model.NetworkInterface{},
		&model.Configuration{},
		&model.Metric{},
		&model.Reports{},
	}

	// if err = db.DropTables(models); err != nil {
	// 	panic(err)
	// }

	if err = db.Migrate(models); err != nil {
		panic(err)
	}

	svc := service.NewService(repository.NewRepository(db.GetDB()))

	// Запускаем симуляцию соединений
	//go simulateConnections(svc)

	// Запускаем симуляцию логов
	//go simulateLogs(svc)

	// Запускаем симуляцию метрик
	//go simulateMetrics(svc)

	// Запускаем симуляцию создания устройств
	//go simulateDeviceCreation(svc)

	//go GenerateDate(db.GetDB())

	app := NewApp()

	err = wails.Run(&options.App{
		Title:  "nws",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			svc,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func GenerateDate(db *gorm.DB) {
	var (
		devices          []model.NetworkDevice
		topologies       []model.NetworkTopology
		topologyElements []model.TopologyElement
		reports          []model.Reports
		//connections      []model.Connection
		//logs             []model.NetworkLog
		//interfaces       []model.NetworkInterface
		//configurations   []model.Configuration
		//metrics          []model.Metric
	)

	// Create Devices
	{
		deviceTypes := []string{"Switch", "Router", "Firewall", "Access Point", "Load Balancer"}
		vendors := []string{"Cisco", "Juniper", "MikroTik", "Ubiquiti", "HP", "Huawei"}
		models := []string{"Catalyst 9300", "MX480", "CCR1036", "EdgeRouter X", "Aruba 2930F", "NE40E"}
		oses := []string{"IOS XE", "JunOS", "RouterOS", "EdgeOS", "ProVision", "VRP"}

		for i := 0; i < 10; i++ {
			device := model.NetworkDevice{
				Hostname:    fmt.Sprintf("%s-%d", deviceTypes[rand.Intn(len(deviceTypes))], i+1),
				IPAddress:   gofakeit.IPv4Address(),
				Type:        deviceTypes[rand.Intn(len(deviceTypes))],
				Vendor:      vendors[rand.Intn(len(vendors))],
				Model:       models[rand.Intn(len(models))],
				Serial:      gofakeit.UUID(),
				OSVersion:   oses[rand.Intn(len(oses))],
				Status:      gofakeit.RandomString([]string{"up", "down"}),
				LastChecked: time.Now(),
				CreatedAt:   time.Now(),
			}
			devices = append(devices, device)
		}
		db.Create(&devices)
	}
	// Create Topologies
	{
		for i := 0; i < 2; i++ {
			topology := model.NetworkTopology{
				Name:        gofakeit.RandomString([]string{"Topology 1", "Topology 2"}),
				Description: gofakeit.Sentence(10),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			topologies = append(topologies, topology)
		}
		db.Create(&topologies)
	}
	// Create Topology Elements
	{
		for i := 0; i < 10; i++ {
			topologyElement := model.TopologyElement{
				TopologyID:  1,
				DeviceID:    gofakeit.RandomInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				X:           gofakeit.Float32(),
				Y:           gofakeit.Float32(),
				CustomStyle: "color: #FF0000; size: 30",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			topologyElements = append(topologyElements, topologyElement)
		}
		db.Create(&topologyElements)
	}

	// CREATE REPORTS
	{
		for i := 0; i < 10; i++ {
			report := model.Reports{
				DeviceID:    gofakeit.IntN(len(devices)),
				Title:       gofakeit.Sentence(10),
				Description: gofakeit.Sentence(10),
				Status:      gofakeit.RandomString([]string{"pending", "in_progress", "completed", "failed"}),
				Priority:    gofakeit.IntN(5),
				Category:    gofakeit.RandomString([]string{"network", "security", "performance"}),
				Tags:        gofakeit.RandomString([]string{"tag1", "tag2"}),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			reports = append(reports, report)
		}
		db.Create(&reports)
	}

	//
	//// Create Connections
	//{
	//	for i := 0; i < 10; i++ {
	//		connection := model.Connection{
	//			SourceDeviceID:      gofakeit.IntN(len(devices)),
	//			DestinationDeviceID: gofakeit.IntN(len(devices)),
	//			SourceInterfaceID:   gofakeit.IntN(len(interfaces)),
	//			Protocol:            gofakeit.RandomString([]string{"TCP", "UDP", "ICMP"}),
	//			Port:                gofakeit.IntN(65535),
	//			Status:              gofakeit.RandomString([]string{"active", "disabled"}),
	//			Latency:             gofakeit.IntN(1000),
	//			PacketLoss:          gofakeit.IntN(100),
	//			PacketsSent:         gofakeit.IntN(1000),
	//			PacketsReceived:     gofakeit.IntN(1000),
	//			PacketsLost:         gofakeit.IntN(100),
	//			PacketsReordered:    gofakeit.IntN(10),
	//			PacketsDuplicated:   gofakeit.IntN(10),
	//			PacketsCorrupted:    gofakeit.IntN(10),
	//			PacketsReassembled:  gofakeit.IntN(10),
	//			CreatedAt:           time.Now(),
	//			UpdatedAt:           time.Now(),
	//		}
	//		connections = append(connections, connection)
	//	}
	//	db.Create(&connections)
	//}
	//
	////
	//// Create Topology Elements
	//{
	//	for i := 0; i < 10; i++ {
	//		topologyElement := model.TopologyElement{
	//			TopologyID:  1,
	//			DeviceID:    devices[gofakeit.IntN(len(devices))].ID,
	//			X:           gofakeit.Float32(),
	//			Y:           gofakeit.Float32(),
	//			CustomStyle: "color: #FF0000; size: 30",
	//			CreatedAt:   time.Now(),
	//			UpdatedAt:   time.Now(),
	//		}
	//		topologyElements = append(topologyElements, topologyElement)
	//	}
	//	db.Create(&topologyElements)
	//}
	//
	//// Create Logs
	//{
	//	for i := 0; i < 10; i++ {
	//		log := model.NetworkLog{
	//			DeviceID:  devices[gofakeit.IntN(len(devices))].ID,
	//			Timestamp: time.Now(),
	//			Action:    gofakeit.RandomString([]string{"ping", "set_config", "get_config"}),
	//			Message:   gofakeit.Sentence(10),
	//			CreatedAt: time.Now(),
	//			UpdatedAt: time.Now(),
	//		}
	//		logs = append(logs, log)
	//	}
	//	db.Create(&logs)
	//}
	//
	//// Create Interfaces
	//{
	//	for i := 0; i < 10; i++ {
	//		networkInterface := model.NetworkInterface{
	//			DeviceID:   devices[gofakeit.IntN(len(devices))].ID,
	//			Name:       gofakeit.Name(),
	//			IPAddress:  gofakeit.IPv4Address(),
	//			Subnet:     gofakeit.IPv4Address(),
	//			Gateway:    gofakeit.IPv4Address(),
	//			MACAddress: gofakeit.MacAddress(),
	//			Speed:      gofakeit.IntN(1000),
	//			Status:     gofakeit.RandomString([]string{"up", "down"}),
	//			CreatedAt:  time.Now(),
	//			UpdatedAt:  time.Now(),
	//		}
	//		interfaces = append(interfaces, networkInterface)
	//	}
	//	db.Create(&interfaces)
	//}
	//
	//// Create Configurations
	//{
	//	for i := 0; i < 10; i++ {
	//		configuration := model.Configuration{
	//			DeviceID:  devices[gofakeit.IntN(len(devices))].ID,
	//			Version:   gofakeit.AppVersion(),
	//			Content:   gofakeit.RandomString([]string{"config", "config2"}),
	//			Changes:   gofakeit.RandomString([]string{"config", "config2"}),
	//			IsActive:  gofakeit.Bool(),
	//			CreatedAt: time.Now(),
	//			UpdatedAt: time.Now(),
	//		}
	//		configurations = append(configurations, configuration)
	//	}
	//	db.Create(&configurations)
	//}
	//
	//// Create Metrics
	//{
	//	for i := 0; i < 10; i++ {
	//		metric := model.Metric{
	//			DeviceID:  devices[gofakeit.IntN(len(devices))].ID,
	//			CPU:       gofakeit.IntN(100),
	//			Memory:    gofakeit.IntN(100),
	//			Network:   gofakeit.IntN(100),
	//			Uptime:    gofakeit.IntN(100),
	//			Interval:  gofakeit.RandomString([]string{"1s", "5s", "10s", "30s", "1m", "5m", "10m", "30m", "1h", "2h", "4h", "8h", "12h", "1d"}),
	//			Status:    gofakeit.RandomString([]string{"up", "down"}),
	//			CreatedAt: time.Now(),
	//			UpdatedAt: time.Now(),
	//		}
	//		metrics = append(metrics, metric)
	//	}
	//	db.Create(&metrics)
	//}

	println("Done")
}
