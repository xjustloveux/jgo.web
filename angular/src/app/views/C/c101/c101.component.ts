import {Component, OnInit} from '@angular/core';
import {AbstractControl, FormGroup, FormBuilder, Validators} from "@angular/forms";
import {Title} from '@angular/platform-browser';
import Swal from 'sweetalert2';
import {C101Service} from './c101.service';
import {VerHsClass} from "../../../model/verHsClass";
import {VerUpdateClass} from "../../../model/verUpdateClass";

@Component({
  selector: 'app-c101',
  templateUrl: './c101.component.html',
  styleUrls: ['./c101.component.css']
})
export class C101Component implements OnInit {

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
  isValid = true;
  verSList: VerHsClass[] = [];
  verEList: VerHsClass[] = [];
  verUpdateList: VerUpdateClass[] = [];
  noneUpdateText = '';

  constructor(
    public service: C101Service,
    public title: Title,
    private fb: FormBuilder
  ) {
    this.title.setTitle('Update | JGo');
    this.form = this.fb.group({
      verS: ['', [Validators.required]],
      verE: ['', [Validators.required]],
    });
  }

  ngOnInit(): void {
    this.service.queryVerList().then((res) => {
      if (res.success) {
        this.verSList = res.list;
      } else {
        Swal.fire('Query version list failed!');
      }
    });
    this.form.get('verS')?.valueChanges.subscribe(selectedValue => {
      this.service.queryVerList(selectedValue).then((res) => {
        if (res.success) {
          this.verEList = res.list;
          this.verUpdateList = [];
          this.form.get('verE')?.setValue('');
        } else {
          Swal.fire('Query version list failed!');
        }
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
    let msg = '';
    if (!ac?.disabled && !ac?.valid && !this.isValid) {
      Object.keys(self.errors).forEach((key: string) => {
        const error = ac?.errors?.[key];
        if (error) {
          const reg = new RegExp("\\{0\\}");
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
      if (res.success) {
        this.verUpdateList = res.list;
        this.noneUpdateText = (this.verUpdateList.length <= 0 ? 'You don\'t need to do anything.' : '');
      } else {
        Swal.fire('Query version update failed!');
      }
    });
  }
}
