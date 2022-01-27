import { Component, ChangeDetectorRef } from '@angular/core';
import { ApiService } from './api/services/api.service';
import { environment } from '../environments/environment';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
    title = 'ui';
    faPlus = faPlus;
    faMinus = faMinus;

    frequency: string
    mode: string
    modeOptions: string[]
    tuningStep: number

    constructor(private apiService: ApiService,
                private changeDetection: ChangeDetectorRef) {}

    ngOnInit(): void {
        console.log("init")
        console.log(environment.apiRootURL);
        this.getFrequency();
        this.getModeOptions();
        this.getMode();
        this.getTuningStep();
    }

    powerOn(): void {
        this.apiService.powerstatPut({"body": {"status": 1}}).subscribe(resp => this.ngOnInit());
    }

    powerOff(): void {
        this.apiService.powerstatPut({"body": {"status": 0}}).subscribe(resp => console.log(resp));
    }

    getFrequency(): void {
        this.apiService.frequencyGet().subscribe(resp => (
            this.frequency = Number(resp.data.frequency).toLocaleString("en-GB") //TODO change to variable locale
        ))
    }

    setFrequency(f: string): void {
        console.log("setFrequency: "+f)
        this.apiService.frequencyPut({"body": {"frequency": f.replace(/,/g, '')}}).subscribe(resp =>
            this.frequency = Number(resp.data.frequency).toLocaleString("en-GB")
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

    getTuningStep(): void {
        this.apiService.tuningStepGet().subscribe(resp => {
            this.tuningStep = resp.data.step
        })
    }

    setTuningStep(s: number): void {
        console.log("Setting Tuning Step " +s)
        this.apiService.tuningStepPut({"body": {"step": s}}).subscribe(resp => this.tuningStep = resp.data.step);
    }

    increaseFrequency(): void {
        this.apiService.frequencyGet().subscribe(resp => {
            this.frequency = Number(resp.data.frequency).toLocaleString("en-GB")
            let _f = Number(this.frequency.replace(/,/g, ''))
            let _new_f = (_f + this.tuningStep).toString(10)
            this.setFrequency(_new_f)
        })
    }

    decreaseFrequency(): void {
        this.apiService.frequencyGet().subscribe(resp => {
            this.frequency = Number(resp.data.frequency).toLocaleString("en-GB")
            let _f = Number(this.frequency.replace(/,/g, ''))
            let _new_f = (_f - this.tuningStep).toString(10)
            this.setFrequency(_new_f)
        })
    }
}
