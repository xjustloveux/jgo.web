import {Component, OnInit} from '@angular/core';
import {MenuService} from './menu.service';
import {MenuClass} from '../../model/menuClass';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {

  menu: MenuClass[] | undefined;
  menuOpen = false;

  constructor(public service: MenuService) {
    this.service.getMenu().then((res) => {
      this.menu = res;
    });
  }

  ngOnInit(): void {
  }

  bnClick(): void {
    this.menuOpen = !this.menuOpen
  }
}
