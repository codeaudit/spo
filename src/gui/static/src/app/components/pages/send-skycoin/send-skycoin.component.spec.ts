import * as sinon from 'sinon';
import { async, inject, ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpModule, XHRBackend } from '@angular/http';
import { MdCardModule, MdSelectModule, MdSnackBarModule, MdTooltipModule } from '@angular/material';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { MockBackend } from '@angular/http/testing';

import { NgxDatatableModule } from '@swimlane/ngx-datatable';
import { ButtonComponent } from '../../layout/button/button.component';
import { SendSkycoinComponent } from './send-skycoin.component';
import { DateTimePipe } from '../../../pipes/date-time.pipe';
import { SkyPipe } from '../../../pipes/sky.pipe';
import { ApiService } from '../../../services/api.service';
import { WalletService } from '../../../services/wallet.service';

describe('SendSkycoinComponent', () => {
  let component: SendSkycoinComponent;
  let mockBackend: MockBackend;
  let fixture: ComponentFixture<SendSkycoinComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        ButtonComponent,
        DateTimePipe,
        SendSkycoinComponent,
        SkyPipe
      ],
      imports: [
        HttpModule,
        MdCardModule,
        MdSelectModule,
        MdSnackBarModule,
        MdTooltipModule,
        NgxDatatableModule,
        NoopAnimationsModule,
        ReactiveFormsModule,
        RouterTestingModule
      ],
      providers: [
        ApiService,
        WalletService,
        { provide: XHRBackend, useClass: MockBackend }
      ]
    })
    .compileComponents();
  }));

  beforeEach(inject([XHRBackend], (backend: MockBackend) => {
    mockBackend = backend;
    fixture = TestBed.createComponent(SendSkycoinComponent);
    component = fixture.componentInstance;
    component.ngOnInit();
    fixture.detectChanges();
  }));

  afterEach(() => {
    fixture.destroy();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('send should not make any backend call if wallet is not chosen', () => {
    const backendCallSpy = sinon.spy();
    mockBackend.connections.subscribe(backendCallSpy);

    component.send();

    sinon.assert.notCalled(backendCallSpy);
  });

  it('send should not make any backend call if address is not set', () => {
    const backendCallSpy = sinon.spy();
    mockBackend.connections.subscribe(backendCallSpy);

    component.form.controls.wallet.setValue({
      meta: { filename: 'test' },
      balance: 1000000
    });

    component.send();

    sinon.assert.notCalled(backendCallSpy);
  });

  it('send should not make any backend call if amount is not set', () => {
    const backendCallSpy = sinon.spy();
    mockBackend.connections.subscribe(backendCallSpy);

    component.form.controls.wallet.setValue({
      meta: { filename: 'test' },
      balance: 1000000
    });
    component.form.controls.address.setValue('0xABC');

    component.send();

    sinon.assert.notCalled(backendCallSpy);
  });

  it('send should not make any backend call if amount exceeds balance', () => {
    const backendCallSpy = sinon.spy();
    mockBackend.connections.subscribe(backendCallSpy);

    component.form.controls.wallet.setValue({
      meta: { filename: 'test' },
      balance: 1000000
    });
    component.form.controls.address.setValue('0xABC');
    component.form.controls.amount.setValue(2);

    component.send();

    sinon.assert.notCalled(backendCallSpy);
  });

  it('send should make backend call if send form is valid', () => {
    let request;
    mockBackend.connections.subscribe((connection) => {
      request = connection.request;
    });

    component.form.controls.wallet.setValue({
      meta: { filename: 'test' },
      balance: 1000000
    });
    component.form.controls.address.setValue('0xABC');
    component.form.controls.amount.setValue(1);

    component.send();

    expect(request).toBeDefined();
    expect(request.url).toBe('http://127.0.0.1:8620/wallet/spend?');
    expect(request.getBody()).toBe('id=test&dst=0xABC&coins=1000000');
  });
});
