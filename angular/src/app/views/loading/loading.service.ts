import {Injectable} from '@angular/core';

import {Subject, Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LoadingService {

  private subject = new Subject<any>();

  constructor() {
  }

  getObservable(): Observable<any> {
    return this.subject.asObservable();
  }

  show() {
    this.subject.next(true);
  }

  hide() {
    this.subject.next(false);
  }
}
