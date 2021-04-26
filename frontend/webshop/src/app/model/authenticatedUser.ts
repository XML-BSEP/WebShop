import { Token } from './token';
import { Role } from './role';
export class AuthenticatedUser {
    public id : number;
    public role : Role;
    public username : string;
    public token : Token;
    public refreshToken : string
    public firstTimeChanged : boolean;
    
    constructor(id : number, role : Role, username : string, token : Token, firstTimeChanged : boolean, rToken : string) {
        this.id = id;
        this.refreshToken = rToken;
        this.role = role;
        this.username = username;
        this.token = token;
        this.firstTimeChanged = firstTimeChanged;
    }
}
