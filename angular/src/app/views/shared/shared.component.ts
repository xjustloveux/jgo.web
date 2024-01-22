import {afterNextRender, Component} from '@angular/core';

@Component({
  selector: 'app-global',
  templateUrl: './shared.component.html',
})
export class SharedComponent {

  // 是否已經render完畢
  isRender: boolean = false;

  constructor() {
    afterNextRender(() => {
      this.isRender = true;
    });
  }

  onActivate(event: any): void {
    if (!this.isRender) {
      return;
    }
    if (typeof window.scroll === 'function') {
      window.scroll(0, 0);
    }
  }
}
