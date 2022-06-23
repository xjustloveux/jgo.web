import {Component, OnInit} from '@angular/core';
import {A101Service} from './a101.service';

@Component({
  selector: 'app-a101',
  templateUrl: './a101.component.html',
  styleUrls: ['./a101.component.css']
})
export class A101Component implements OnInit {

  constructor(
    public service: A101Service,
  ) {
  }

  ngOnInit(): void {
  }
}
