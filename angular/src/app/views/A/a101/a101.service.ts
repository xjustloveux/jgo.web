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
}
