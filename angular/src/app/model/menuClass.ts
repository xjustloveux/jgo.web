export class MenuClass {
  PRO_URL: string;
  PRO_NAME: string;
  LIST: MenuClass[];

  constructor(data: any) {
    this.PRO_URL = data.PRO_URL;
    this.PRO_NAME = data.PRO_NAME;
    this.LIST = data.LIST;
  }
}
