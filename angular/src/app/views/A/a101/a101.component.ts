import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
import {A101Service} from './a101.service';

@Component({
  selector: 'app-a101',
  templateUrl: './a101.component.html',
  styleUrls: ['./a101.component.css']
})
export class A101Component implements OnInit {

  apiList = [
    {name: 'jcast', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jcast'},
    {name: 'jconf', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jconf'},
    {name: 'jcron', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jcron'},
    {name: 'jevent', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jevent'},
    {name: 'jfile', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jfile'},
    {name: 'jlog', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jlog'},
    {name: 'jruntime', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jruntime'},
    {name: 'jslice', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jslice'},
    {name: 'jsql', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jsql'},
    {name: 'jtime', url: 'https://pkg.go.dev/github.com/xjustloveux/jgo/jtime'},
  ];

  constructor(
    public service: A101Service,
    public title: Title
  ) {
    this.title.setTitle('JGo | Golang');
  }

  ngOnInit(): void {
  }
}
