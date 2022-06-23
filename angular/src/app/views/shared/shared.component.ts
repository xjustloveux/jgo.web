import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-global',
  templateUrl: './shared.component.html',
  styleUrls: ['./shared.component.css']
})
export class SharedComponent implements OnInit {

  constructor() {
  }

  ngOnInit(): void {
  }

  onActivate(event: any) {
    if (typeof window.scroll === 'function') {
      window.scroll(0, 0);
    }
  }
}
