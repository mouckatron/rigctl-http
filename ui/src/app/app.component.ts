import { Component, ChangeDetectorRef } from '@angular/core';
import { ApiService } from './api/services/api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
    title = 'ui';

    frequency: string
    mode: string
    modeOptions: string[]

    constructor(private apiService: ApiService,
                private changeDetection: ChangeDetectorRef) {}

    ngOnInit(): void {
        console.log("init")
        this.getFrequency();
        this.getModeOptions();
        this.getMode();
    }

    powerOn(): void {
        this.apiService.powerstatPut({"body": {"status": 1}}).subscribe(resp => this.ngOnInit());
    }

    powerOff(): void {
        this.apiService.powerstatPut({"body": {"status": 0}}).subscribe(resp => console.log(resp));
    }

    getFrequency(): void {
        this.apiService.frequencyGet().subscribe(resp => (
            this.frequency = resp.data.frequency
        ))
    }

    setFrequency(f: string): void {
        this.apiService.frequencyPut({"body": {"frequency": f}}).subscribe(resp =>
            this.frequency = resp.data.frequency
        );
    }

    getMode(): void {
        this.apiService.modeGet().subscribe(resp => this.mode = resp.data.mode)
    }

    getModeOptions(): void {
        this.apiService.modeOptionsGet().subscribe(resp => this.modeOptions = resp.data.options)
    }

    setMode(m: string): void {
        console.log("Setting mode " +m)
        this.apiService.modePut({"body": {"mode": m}}).subscribe(resp => {
            if(!resp.success) this.modeOptions = resp.data.options
        })
    }

}
