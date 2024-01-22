import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

import Swal from 'sweetalert2';

import {LoadingService} from '../views/loading/loading.service';
import {Message} from "../const/message";

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(
    private http: HttpClient,
    private loadingService: LoadingService
  ) {
  }

  /**
   * get
   * @param url      api url
   * @param json     json object
   * @param loading loading screen
   */
  get(url: string, json: any, loading?: boolean): Promise<any> {
    return new Promise((resolve, reject) => {

      if (loading !== false) {
        this.loadingService.show();
      }

      let params: string = '';
      for (const k in json) {
        const val = json[k];
        if (val) {
          params += (params === '' ? '?' : '&') + k + '=' + encodeURIComponent(val);
        }
      }

      this.http
        .get(url + params)
        .subscribe({
          next: (res: any) => {
            if (res.success) {
              resolve(res.data);
            } else {
              Swal.fire(Message[res.code]);
              reject(res.code);
            }
          },
          error: (error) => {
            console.log(error);
            reject(error);
            Swal.fire('Connection failed or timed out!');
          },
          complete: () => {
            this.loadingService.hide();
          },
        });
    });
  }

  /**
   * post
   * @param url      api url
   * @param json     json object
   * @param loading loading screen
   */
  post(url: string, json: any, loading?: boolean): Promise<any> {
    return new Promise((resolve, reject) => {

      if (loading !== false) {
        this.loadingService.show();
      }

      this.http
        .post(url, json)
        .subscribe({
          next: (res: any) => {
            if (res.success) {
              resolve(res.data);
            } else {
              Swal.fire(Message[res.code]);
              reject(res.code);
            }
          },
          error: (error) => {
            console.log(error);
            reject(error);
            Swal.fire('Connection failed or timed out!');
          },
          complete: () => {
            this.loadingService.hide();
          },
        });
    });
  }
}
