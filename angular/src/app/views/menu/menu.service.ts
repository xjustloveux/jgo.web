import {Injectable} from '@angular/core';
import {HttpService} from '../../shared/http.service';

@Injectable({
  providedIn: 'root'
})
export class MenuService {

  constructor(
    private http: HttpService
  ) {
  }

  getMenu(): Promise<any> {
    return new Promise((resolve, reject) => {
      this.http.get('/Shared/QueryMenu', true).then((res: any) => {
        resolve(res);
      });
    });
  }
}
