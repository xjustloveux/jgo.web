export class VerUpdateClass {
  seq: number;
  ver_s: number;
  ver_e: number;
  content: string;

  constructor(data: any) {
    this.seq = data.seq;
    this.ver_s = data.ver_s;
    this.ver_e = data.ver_e;
    this.content = data.content;
  }
}
