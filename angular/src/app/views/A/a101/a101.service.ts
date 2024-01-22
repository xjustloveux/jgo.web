import {Injectable} from '@angular/core';

import {HttpService} from '../../../shared/http.service';

@Injectable({
  providedIn: 'root'
})
export class A101Service {

  constructor(
    private http: HttpService
  ) {
  }

  queryPage(data: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/a/a101', data, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  createMsg(data: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.post('/a/a101/CreateMsg', data, true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
