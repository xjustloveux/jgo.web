import {Component, OnInit} from '@angular/core';
import {Subscription} from 'rxjs';
import {LoadingService} from './loading.service';

@Component({
  selector: 'app-loading',
  templateUrl: './loading.component.html',
  styleUrls: ['./loading.component.css']
})
export class LoadingComponent implements OnInit {

  show = false;
  subscription: Subscription = new Subscription;

  constructor(private loadingService: LoadingService) {
  }

  ngOnInit() {
    this.subscription = this.loadingService.getObservable().subscribe((show) => {
      this.show = show;
    });
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
