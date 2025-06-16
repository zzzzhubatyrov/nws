export namespace models {
	
	export class Alert {
	    id: number;
	    device: string;
	    type: string;
	    severity: string;
	    time: string;
	    status: string;
	    title: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Alert(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.device = source["device"];
	        this.type = source["type"];
	        this.severity = source["severity"];
	        this.time = source["time"];
	        this.status = source["status"];
	        this.title = source["title"];
	        this.message = source["message"];
	    }
	}
	export class Connection {
	    id: number;
	    source: number;
	    target: number;
	    status: string;
	    bandwidth: string;
	    latency: number;
	    type: string;
	    name: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.source = source["source"];
	        this.target = source["target"];
	        this.status = source["status"];
	        this.bandwidth = source["bandwidth"];
	        this.latency = source["latency"];
	        this.type = source["type"];
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class Device {
	    // Go type: time
	    createdAt: any;
	    id: number;
	    name: string;
	    type: string;
	    ip: string;
	    status: string;
	    cpu: number;
	    memory: number;
	    uptime: string;
	    x: number;
	    y: number;
	    vendor: string;
	    model: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.ip = source["ip"];
	        this.status = source["status"];
	        this.cpu = source["cpu"];
	        this.memory = source["memory"];
	        this.uptime = source["uptime"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.vendor = source["vendor"];
	        this.model = source["model"];
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
	export class Log {
	    id: number;
	    deviceId?: number;
	    connectionId?: number;
	    level: string;
	    message: string;
	    // Go type: time
	    time: any;
	
	    static createFrom(source: any = {}) {
	        return new Log(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.deviceId = source["deviceId"];
	        this.connectionId = source["connectionId"];
	        this.level = source["level"];
	        this.message = source["message"];
	        this.time = this.convertValues(source["time"], null);
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
	    id: number;
	    deviceId?: number;
	    connectionId?: number;
	    name: string;
	    value: number;
	    // Go type: time
	    time: any;
	
	    static createFrom(source: any = {}) {
	        return new Metric(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.deviceId = source["deviceId"];
	        this.connectionId = source["connectionId"];
	        this.name = source["name"];
	        this.value = source["value"];
	        this.time = this.convertValues(source["time"], null);
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
	    id: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    email: string;
	    name: string;
	    role: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.email = source["email"];
	        this.name = source["name"];
	        this.role = source["role"];
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

export namespace service {
	
	export class TokenClaims {
	    user_id: number;
	    email: string;
	    role: string;
	    iss?: string;
	    sub?: string;
	    aud?: string[];
	    // Go type: jwt
	    exp?: any;
	    // Go type: jwt
	    nbf?: any;
	    // Go type: jwt
	    iat?: any;
	    jti?: string;
	
	    static createFrom(source: any = {}) {
	        return new TokenClaims(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.user_id = source["user_id"];
	        this.email = source["email"];
	        this.role = source["role"];
	        this.iss = source["iss"];
	        this.sub = source["sub"];
	        this.aud = source["aud"];
	        this.exp = this.convertValues(source["exp"], null);
	        this.nbf = this.convertValues(source["nbf"], null);
	        this.iat = this.convertValues(source["iat"], null);
	        this.jti = source["jti"];
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

