export namespace models {
	
	export class Usuario {
	    id: number;
	    name: string;
	    last_name: string;
	    username: string;
	    password: string;
	    porfile_image: string;
	
	    static createFrom(source: any = {}) {
	        return new Usuario(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.last_name = source["last_name"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.porfile_image = source["porfile_image"];
	    }
	}

}

