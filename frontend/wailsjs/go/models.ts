export namespace gorm {
	
	export class DeletedAt {
	    Time: time.Time;
	    Valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DeletedAt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Time = this.convertValues(source["Time"], time.Time);
	        this.Valid = source["Valid"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace model {
	
	export class NetworkLog {
	    ID: number;
	    device_id: number;
	    device: NetworkDevice;
	    timestamp: time.Time;
	    action: string;
	    message: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new NetworkLog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.timestamp = this.convertValues(source["timestamp"], time.Time);
	        this.action = source["action"];
	        this.message = source["message"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NetworkInterface {
	    ID: number;
	    device_id: number;
	    device: NetworkDevice;
	    name: string;
	    ip_address: string;
	    subnet: string;
	    gateway: string;
	    mac_address: string;
	    speed: number;
	    status: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new NetworkInterface(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.name = source["name"];
	        this.ip_address = source["ip_address"];
	        this.subnet = source["subnet"];
	        this.gateway = source["gateway"];
	        this.mac_address = source["mac_address"];
	        this.speed = source["speed"];
	        this.status = source["status"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Metric {
	    ID: number;
	    device_id: number;
	    device: NetworkDevice;
	    cpu: number;
	    memory: number;
	    network: number;
	    uptime: number;
	    interval: string;
	    status: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new Metric(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.cpu = source["cpu"];
	        this.memory = source["memory"];
	        this.network = source["network"];
	        this.uptime = source["uptime"];
	        this.interval = source["interval"];
	        this.status = source["status"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Connection {
	    ID: number;
	    source_device_ip: string;
	    source_device: NetworkDevice;
	    destination_device_ip: string;
	    destination_device: NetworkDevice;
	    protocol: string;
	    port: number;
	    status: string;
	    latency: number;
	    packet_loss: number;
	    packets_sent: number;
	    packets_received: number;
	    packets_lost: number;
	    packets_reordered: number;
	    packets_duplicated: number;
	    packets_corrupted: number;
	    packets_reassembled: number;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.source_device_ip = source["source_device_ip"];
	        this.source_device = this.convertValues(source["source_device"], NetworkDevice);
	        this.destination_device_ip = source["destination_device_ip"];
	        this.destination_device = this.convertValues(source["destination_device"], NetworkDevice);
	        this.protocol = source["protocol"];
	        this.port = source["port"];
	        this.status = source["status"];
	        this.latency = source["latency"];
	        this.packet_loss = source["packet_loss"];
	        this.packets_sent = source["packets_sent"];
	        this.packets_received = source["packets_received"];
	        this.packets_lost = source["packets_lost"];
	        this.packets_reordered = source["packets_reordered"];
	        this.packets_duplicated = source["packets_duplicated"];
	        this.packets_corrupted = source["packets_corrupted"];
	        this.packets_reassembled = source["packets_reassembled"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NetworkDevice {
	    ID: number;
	    hostname: string;
	    ip_address: string;
	    type: string;
	    vendor: string;
	    model: string;
	    serial: string;
	    os_version: string;
	    status: string;
	    last_checked: time.Time;
	    connections: Connection[];
	    metrics: Metric[];
	    network_interfaces: NetworkInterface[];
	    network_logs: NetworkLog[];
	    configurations: Configuration[];
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new NetworkDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.hostname = source["hostname"];
	        this.ip_address = source["ip_address"];
	        this.type = source["type"];
	        this.vendor = source["vendor"];
	        this.model = source["model"];
	        this.serial = source["serial"];
	        this.os_version = source["os_version"];
	        this.status = source["status"];
	        this.last_checked = this.convertValues(source["last_checked"], time.Time);
	        this.connections = this.convertValues(source["connections"], Connection);
	        this.metrics = this.convertValues(source["metrics"], Metric);
	        this.network_interfaces = this.convertValues(source["network_interfaces"], NetworkInterface);
	        this.network_logs = this.convertValues(source["network_logs"], NetworkLog);
	        this.configurations = this.convertValues(source["configurations"], Configuration);
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Configuration {
	    ID: number;
	    device_id: number;
	    device: NetworkDevice;
	    version: string;
	    content: string;
	    changes: string;
	    is_active: boolean;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new Configuration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.version = source["version"];
	        this.content = source["content"];
	        this.changes = source["changes"];
	        this.is_active = source["is_active"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class User {
	    ID: number;
	    username: string;
	    password: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LoginResponse {
	    user?: User;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new LoginResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.user = this.convertValues(source["user"], User);
	        this.token = source["token"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	export class NetworkTopology {
	    ID: number;
	    name: string;
	    description: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new NetworkTopology(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Reports {
	    ID: number;
	    device_id: number;
	    device: NetworkDevice;
	    title: string;
	    description: string;
	    status: string;
	    priority: number;
	    category: string;
	    tags: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new Reports(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.priority = source["priority"];
	        this.category = source["category"];
	        this.tags = source["tags"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TopologyElement {
	    ID: number;
	    topology_id: number;
	    topology: NetworkTopology;
	    device_id: number;
	    device: NetworkDevice;
	    x: number;
	    y: number;
	    custom_style: string;
	    CreatedAt: time.Time;
	    UpdatedAt: time.Time;
	    DeletedAt: gorm.DeletedAt;
	
	    static createFrom(source: any = {}) {
	        return new TopologyElement(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.topology_id = source["topology_id"];
	        this.topology = this.convertValues(source["topology"], NetworkTopology);
	        this.device_id = source["device_id"];
	        this.device = this.convertValues(source["device"], NetworkDevice);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.custom_style = source["custom_style"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], time.Time);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], time.Time);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], gorm.DeletedAt);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace time {
	
	export class Time {
	
	
	    static createFrom(source: any = {}) {
	        return new Time(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

