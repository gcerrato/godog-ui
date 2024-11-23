export namespace fs {
	
	export class FileScanner {
	
	
	    static createFrom(source: any = {}) {
	        return new FileScanner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

