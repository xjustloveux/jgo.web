import {Injectable} from '@angular/core';

import {HttpService} from '../../../shared/http.service';

@Injectable({
  providedIn: 'root'
})
export class C101Service {

  constructor(
    private http: HttpService
  ) {
  }

  queryVerList(ver?: number): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/c/c101/QueryVer', { ver: ver }, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  queryVerUpdate(data: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/c/c101/QueryVerUpdate', data, true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
