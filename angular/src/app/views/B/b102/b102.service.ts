import {Injectable} from '@angular/core';
import {HttpService} from '../../../shared/http.service';

@Injectable({
  providedIn: 'root'
})
export class B102Service {

  constructor(
    private http: HttpService
  ) {
  }

  queryContent(id: string): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/b/' + id, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  queryLog(name: string): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/b/b102/QueryLog?name=' + name, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  trigger(data: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.post('/b/b102/Trigger', data, true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
