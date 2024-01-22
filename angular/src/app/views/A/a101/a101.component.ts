import {afterNextRender, Component} from '@angular/core';
import {Meta, Title} from '@angular/platform-browser';
import {AbstractControl, FormBuilder, FormGroup, Validators} from '@angular/forms';

import Swal from 'sweetalert2';

import {A101Service} from './a101.service';
import {ResMsgClass} from '../../../model/resMsgClass';

@Component({
  selector: 'app-a101',
  templateUrl: './a101.component.html',
})
export class A101Component {

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
  tabs = [
    {active: true, dbType: 'MySql', page: 1, size: 5, total: 0, res: ResMsgClass},
    {active: false, dbType: 'MSSql', page: 1, size: 5, total: 0, res: ResMsgClass},
    {active: false, dbType: 'Oracle', page: 1, size: 5, total: 0, res: ResMsgClass},
    {active: false, dbType: 'PostgreSql', page: 1, size: 5, total: 0, res: ResMsgClass},
  ];
  size: number[] = [5, 10, 15];
  nowSize: number = 5;
  nowPage: number = 1;
  form: FormGroup;
  errors = {
    min: 'Minimum Number {0}',
    max: 'Maximum Number {0}',
    required: 'Required',
    email: 'Email Format Error',
    minlength: 'Minimum Length {0}',
    maxlength: 'Maximum Length {0}',
    pattern: 'Please enter information in valid formats.'
  };
  isValid: boolean = true;

  constructor(
    public service: A101Service,
    private meta: Meta,
    public title: Title,
    private fb: FormBuilder
  ) {
    const ogTitle: string = 'JGo | Golang';
    const description: string = 'JGo provides an easier configuration for writing sql, log, and cron jobs.';
    const keywords: string = 'jgo,golang,cron,sql,log,job,schedule,mssql,mysql,oracle';
    const ogUrl: string = 'https://jgo.dev/home';
    const metaList: any = [
      {name: 'description', content: description},
      {name: 'keywords', content: keywords},
      {name: 'twitter:title', content: ogTitle},
      {name: 'twitter:description', content: description},
      {property: 'og:title', content: ogTitle},
      {property: 'og:description', content: description},
      {property: 'og:url', content: ogUrl},
    ];
    for (const item of metaList) {
      this.meta.updateTag(item);
    }
    this.title.setTitle(ogTitle);
    this.form = this.fb.group({
      content: ['', [Validators.required, Validators.maxLength(50)]],
    });
    afterNextRender(() => {
      this.tabs.forEach(tab => {
        this.queryPage(tab);
      });
    })
  }

  isInvalid(name: string): string {
    const ac: AbstractControl | null = this.form.get(name);
    if (!ac?.disabled && !ac?.valid && !this.isValid) {
      return 'is-invalid';
    }
    return '';
  }

  errorMsg(name: string): string {
    const self: any = this;
    const ac: AbstractControl | null = this.form.get(name);
    let msg = '';
    if (!ac?.disabled && !ac?.valid && !this.isValid) {
      Object.keys(self.errors).forEach((key: string) => {
        const error = ac?.errors?.[key];
        if (error) {
          const reg = new RegExp('\\{0\\}');
          switch (key) {
            case 'min':
            case 'max': {
              const num = error[key];
              msg = self.errors[key].replace(reg, num);
            }
              break;
            case 'minlength':
            case 'maxlength': {
              const num = error.requiredLength;
              msg = self.errors[key].replace(reg, num);
            }
              break;
            default: {
              msg = self.errors[key];
            }
              break;
          }
        }
      });
      msg = (msg === '' ? self.errors.pattern : msg);
    }
    return msg;
  }

  queryPage(tab: any): void {
    this.service.queryPage({
      dbType: tab.dbType,
      page: tab.page,
      size: tab.size,
    }).then(res => {
      tab.res = res;
    });
  }

  tabClick(idx: number): void {
    this.tabs.forEach((tab, i) => {
      tab.active = (i == idx);
      if (tab.active) {
        this.nowSize = tab.size;
        this.nowPage = tab.page;
      }
    });
  }

  getActiveTab(): any {
    let t: any;
    this.tabs.forEach((tab: any) => {
      if (tab.active) {
        t = tab;
      }
    });
    return t;
  }

  getActiveTabRes(): ResMsgClass {
    return this.getActiveTab().res;
  }

  getTime(date: Date): string {
    const d = new Date(date);
    return d.getFullYear() + '-' +
      (d.getMonth()+1).toString().padStart(2, '0') + '-' +
      d.getDate().toString().padStart(2, '0') + ' ' +
      d.getHours().toString().padStart(2, '0') + ':' +
      d.getMinutes().toString().padStart(2, '0') + ':' +
      d.getSeconds().toString().padStart(2, '0');
  }

  createMsg(): void {
    this.isValid = this.form.valid;
    if (!this.isValid) {
      Swal.fire('Please fill required fields and correct errors.');
      return;
    }
    const data = this.form.getRawValue();
    let t: any;
    this.tabs.forEach((tab: any) => {
      if (tab.active) {
        data.dbType = tab.dbType;
        t = tab;
      }
    });
    this.service.createMsg(data).then((res) => {
      this.queryPage(t);
    });
  }

  sizeChange(val: number): void {
    const tab = this.getActiveTab();
    this.nowSize = val;
    tab.size = val;
    this.queryPage(tab);
  }

  pageChange(val: number): void {
    if (val < 1) {
      this.nowPage = 1;
    } else {
      this.nowPage = val;
    }
  }

  pageClick(): void {
    const tab = this.getActiveTab();
    tab.page = this.nowPage;
    this.queryPage(tab);
  }

  prevPage(): void {
    const tab = this.getActiveTab();
    tab.page--;
    if (tab.page < 1) {
      tab.page = 1;
    }
    this.nowPage = tab.page;
    this.queryPage(tab);
  }

  nextPage(): void {
    const tab = this.getActiveTab();
    tab.page++;
    this.nowPage = tab.page;
    this.queryPage(tab);
  }
}
