import {afterNextRender, Component} from '@angular/core';
import {AbstractControl, FormGroup, FormBuilder, Validators} from '@angular/forms';
import {Meta, Title} from '@angular/platform-browser';

import Swal from 'sweetalert2';

import {C101Service} from './c101.service';
import {VerHsClass} from '../../../model/verHsClass';
import {VerUpdateClass} from '../../../model/verUpdateClass';

@Component({
  selector: 'app-c101',
  templateUrl: './c101.component.html',
})
export class C101Component {

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
  verSList: VerHsClass[] = [];
  verEList: VerHsClass[] = [];
  verUpdateList: VerUpdateClass[] = [];
  noneUpdateText: string = '';

  constructor(
    public service: C101Service,
    private meta: Meta,
    public title: Title,
    private fb: FormBuilder
  ) {
    const ogTitle: string = 'Update | JGo';
    const description: string = 'JGo version upgrade steps.';
    const keywords: string = 'jgo,update,version,upgrade';
    const ogUrl: string = 'https://jgo.dev/pkg/update';
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
      verS: ['', [Validators.required]],
      verE: ['', [Validators.required]],
    });
    afterNextRender(() => {
      this.service.queryVerList().then((res) => {
        this.verSList = res;
      });
      this.form.get('verS')?.valueChanges.subscribe(selectedValue => {
        this.service.queryVerList(selectedValue).then((res) => {
          this.verEList = res;
          this.verUpdateList = [];
          this.form.get('verE')?.setValue('');
        });
      });
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

  query(): void {
    this.isValid = this.form.valid;
    if (!this.isValid) {
      Swal.fire('Please fill required fields and correct errors.');
      return;
    }
    this.service.queryVerUpdate(this.form.getRawValue()).then((res) => {
      this.verUpdateList = res;
      this.noneUpdateText = (this.verUpdateList.length <= 0 ? 'You don\'t need to do anything.' : '');
    });
  }
}
