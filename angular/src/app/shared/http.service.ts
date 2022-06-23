import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import Swal from 'sweetalert2';
import {LoadingService} from '../views/loading/loading.service';

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
   * @param loading loading screen
   */
  get(url: string, loading?: boolean): Promise<any> {
    return new Promise((resolve, reject) => {

      if (loading !== false) {
        this.loadingService.show();
      }

      this.http
        .get(url)
        .subscribe({
          next: (res) => {
            resolve(res);
          },
          error: (error) => {
            console.log(error);
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
          next: (res) => {
            resolve(res);
          },
          error: (error) => {
            console.log(error);
            Swal.fire('Connection failed or timed out!');
          },
          complete: () => {
            this.loadingService.hide();
          },
        });
    });
  }
}
