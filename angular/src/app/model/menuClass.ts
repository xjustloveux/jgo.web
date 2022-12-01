export class MenuClass {
  url: string;
  name: string;
  child: MenuClass[];

  constructor(data: any) {
    this.url = data.url;
    this.name = data.name;
    this.child = data.child;
  }
}
