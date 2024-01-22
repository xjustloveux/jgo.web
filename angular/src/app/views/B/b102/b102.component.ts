import {afterNextRender, Component} from '@angular/core';
import {Meta, Title} from '@angular/platform-browser';
import {AbstractControl, FormBuilder, FormGroup, Validators} from '@angular/forms';

import Swal from 'sweetalert2';

import {B102Service} from './b102.service';
import {PkgContentClass} from '../../../model/pkgContentClass';
import {JobLogClass} from '../../../model/jobLogClass';

@Component({
  selector: 'app-b102',
  templateUrl: './b102.component.html',
})
export class B102Component {

  proId: string = 'b102';
  content: PkgContentClass[] = [];
  log: { job001: JobLogClass[], job002: JobLogClass[] } = {job001: [], job002: []};
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
    public service: B102Service,
    private meta: Meta,
    public title: Title,
    private fb: FormBuilder
  ) {
    const ogTitle: string = 'jCron | JGo';
    const description: string = 'jCron is a cron job scheduler written in Go.';
    const keywords: string = 'jcron,cron,job,schedule';
    const ogUrl: string = 'https://jgo.dev/pkg/jcron';
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
      id: ['', [Validators.required, Validators.maxLength(10)]],
    });
    afterNextRender(() => {
      this.service.queryContent(this.proId).then((res) => {
        this.content = res;
      });
      this.queryLog('job001');
      this.queryLog('job002');
    });
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
    let msg: string = '';
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

  queryLog(name: 'job001' | 'job002'): void {
    this.service.queryLog(name).then((res) => {
      this.log[name] = res;
    });
  }

  trigger(): void {
    this.isValid = this.form.valid;
    if (!this.isValid) {
      Swal.fire('Please fill required fields and correct errors.');
      return;
    }
    this.service.trigger(this.form.getRawValue()).then((res) => {
      this.queryLog('job002');
    });
  }
}
