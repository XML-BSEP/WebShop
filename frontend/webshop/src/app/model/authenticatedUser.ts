import { Token } from './token';
import { Role } from './role';
export class AuthenticatedUser {
    public id : number;
    public role : Role;
    public access_token: string;
    public refresh_token : string;
    constructor(id : number, role : Role, access_token : string, refresh_token : string) {
        this.id = id;
        this.role = role;
        this.access_token = access_token;
        this.refresh_token = refresh_token;
    }
}
