import {Injectable} from '@angular/core';
import {HttpService} from '../../../shared/http.service';

@Injectable({
  providedIn: 'root'
})
export class B103Service {

  constructor(
    private http: HttpService
  ) {
  }

  queryContent(id: string): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/b/' + id, {}, true).then((res: any) => {
        resolve(res);
      });
    });
  }

  queryLog(): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/b/b103/QueryLog', {}, true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
