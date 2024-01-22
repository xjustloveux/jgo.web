import {afterNextRender, Component} from '@angular/core';

import {MenuService} from './menu.service';
import {MenuClass} from '../../model/menuClass';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
})
export class MenuComponent {

  menu: MenuClass[] | undefined;
  menuOpen: boolean = false;

  constructor(public service: MenuService) {
    afterNextRender(() => {
      this.service.getMenu().then((res) => {
        this.menu = res;
      });
    });
  }

  bnClick(): void {
    this.menuOpen = !this.menuOpen
  }
}
