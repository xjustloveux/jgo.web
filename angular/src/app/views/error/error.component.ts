import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-error',
  templateUrl: './error.component.html',
  styleUrls: ['./error.component.css']
})
export class ErrorComponent implements OnInit {

  constructor(
    public title: Title,
    private router: Router
  ) {
    this.title.setTitle('Error | JGo');
  }

  ngOnInit(): void {
  }

  bnClick(): void {
    this.router.navigate(['/']).then(() => {
    });
  }
}
