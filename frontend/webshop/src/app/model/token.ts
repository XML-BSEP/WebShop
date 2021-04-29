export class Token {
    public access_token: string;
    public refresh_token : string;
    constructor(access_token : string, refresh : string) {
      this.access_token = access_token;
      this.refresh_token = this.refresh_token;
    }

    getToken() {
        return this.access_token;
    }


}
