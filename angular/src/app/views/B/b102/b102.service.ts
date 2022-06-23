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
      this.http.get('/' + id + '/QueryContent', true).then((res: any) => {
        resolve(res);
      });
    });
  }

  queryLog(name: string): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.post('/B102/QueryLog', {NAME: name}, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  trigger(data: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.post('/B102/Trigger', data, true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
