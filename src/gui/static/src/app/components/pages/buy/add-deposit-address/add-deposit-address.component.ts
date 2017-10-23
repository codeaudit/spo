import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WalletService } from '../../../../services/wallet.service';
import { PurchaseService } from '../../../../services/purchase.service';
import { MdDialogRef } from '@angular/material';

@Component({
  selector: 'app-add-deposit-address',
  templateUrl: './add-deposit-address.component.html',
  styleUrls: ['./add-deposit-address.component.css']
})
export class AddDepositAddressComponent implements OnInit {

  form: FormGroup;
  tokenType: string;
  loading: Boolean = false;
  constructor(
    public walletService: WalletService,
    private dialogRef: MdDialogRef<AddDepositAddressComponent>,
    private formBuilder: FormBuilder,
    private purchaseService: PurchaseService,
  ) {}

  ngOnInit() {
    this.initForm();
  }

  generate() {
    this.loading = true;
    this.purchaseService.generate(this.form.value.address, this.tokenType).subscribe(() => {
      this.dialogRef.close();
    }, (err)=>{
      alert(err);  
    }, ()=>{
      this.loading = false;
    });
  }

  private initForm() {
    this.form = this.formBuilder.group({
      address: ['', Validators.required],
    });
  }
}
