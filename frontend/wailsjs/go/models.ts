export namespace models {
	
	export class Usuario {
	    id: number;
	    nombre: string;
	    apellidos: string;
	    email: string;
	    usuario: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Usuario(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nombre = source["nombre"];
	        this.apellidos = source["apellidos"];
	        this.email = source["email"];
	        this.usuario = source["usuario"];
	        this.password = source["password"];
	    }
	}

}

