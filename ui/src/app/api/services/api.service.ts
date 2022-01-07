/* tslint:disable */
/* eslint-disable */
import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { BaseService } from '../base-service';
import { ApiConfiguration } from '../api-configuration';
import { StrictHttpResponse } from '../strict-http-response';
import { RequestBuilder } from '../request-builder';
import { Observable } from 'rxjs';
import { map, filter } from 'rxjs/operators';

import { Frequency } from '../models/frequency';
import { Mode } from '../models/mode';
import { Powerstat } from '../models/powerstat';
import { CmdOptions } from '../models/cmd-options';

@Injectable({
  providedIn: 'root',
})
export class ApiService extends BaseService {
  constructor(
    config: ApiConfiguration,
    http: HttpClient
  ) {
    super(config, http);
  }

  /**
   * Path part for operation frequencyGet
   */
  static readonly FrequencyGetPath = '/frequency';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `frequencyGet()` instead.
   *
   * This method doesn't expect any request body.
   */
  frequencyGet$Response(params?: {
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.FrequencyGetPath, 'get');
    if (params) {
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: Frequency;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `frequencyGet$Response()` instead.
   *
   * This method doesn't expect any request body.
   */
  frequencyGet(params?: {
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}> {

    return this.frequencyGet$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
})
    );
  }

  /**
   * Path part for operation frequencyPut
   */
  static readonly FrequencyPutPath = '/frequency';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `frequencyPut()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  frequencyPut$Response(params?: {
    body?: any
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.FrequencyPutPath, 'put');
    if (params) {
      rb.body(params.body, 'application/json');
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: Frequency;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `frequencyPut$Response()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  frequencyPut(params?: {
    body?: any
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}> {

    return this.frequencyPut$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Frequency;
})
    );
  }

  /**
   * Path part for operation powerstatGet
   */
  static readonly PowerstatGetPath = '/powerstat';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `powerstatGet()` instead.
   *
   * This method doesn't expect any request body.
   */
  powerstatGet$Response(params?: {
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.PowerstatGetPath, 'get');
    if (params) {
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: Powerstat;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `powerstatGet$Response()` instead.
   *
   * This method doesn't expect any request body.
   */
  powerstatGet(params?: {
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}> {

    return this.powerstatGet$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
})
    );
  }

  /**
   * Path part for operation powerstatPut
   */
  static readonly PowerstatPutPath = '/powerstat';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `powerstatPut()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  powerstatPut$Response(params?: {
    body?: any
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.PowerstatPutPath, 'put');
    if (params) {
      rb.body(params.body, 'application/json');
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: Powerstat;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `powerstatPut$Response()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  powerstatPut(params?: {
    body?: any
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}> {

    return this.powerstatPut$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Powerstat;
})
    );
  }

  /**
   * Path part for operation modeGet
   */
  static readonly ModeGetPath = '/mode';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `modeGet()` instead.
   *
   * This method doesn't expect any request body.
   */
  modeGet$Response(params?: {
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Mode;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.ModeGetPath, 'get');
    if (params) {
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: Mode;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `modeGet$Response()` instead.
   *
   * This method doesn't expect any request body.
   */
  modeGet(params?: {
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Mode;
}> {

    return this.modeGet$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Mode;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: Mode;
})
    );
  }

  /**
   * Path part for operation modePut
   */
  static readonly ModePutPath = '/mode';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `modePut()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  modePut$Response(params?: {
    body?: any
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: any;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.ModePutPath, 'put');
    if (params) {
      rb.body(params.body, 'application/json');
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: any;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `modePut$Response()` instead.
   *
   * This method sends `application/json` and handles request body of type `application/json`.
   */
  modePut(params?: {
    body?: any
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: any;
}> {

    return this.modePut$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: any;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: any;
})
    );
  }

  /**
   * Path part for operation modeOptionsGet
   */
  static readonly ModeOptionsGetPath = '/mode/_options';

  /**
   * This method provides access to the full `HttpResponse`, allowing access to response headers.
   * To access only the response body, use `modeOptionsGet()` instead.
   *
   * This method doesn't expect any request body.
   */
  modeOptionsGet$Response(params?: {
  }): Observable<StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: CmdOptions;
}>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.ModeOptionsGetPath, 'get');
    if (params) {
    }

    return this.http.request(rb.build({
      responseType: 'json',
      accept: 'application/json'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return r as StrictHttpResponse<{
        'success'?: boolean;
        'error'?: string;
        'raw'?: string;
        'data'?: CmdOptions;
        }>;
      })
    );
  }

  /**
   * This method provides access to only to the response body.
   * To access the full response (for headers, for example), `modeOptionsGet$Response()` instead.
   *
   * This method doesn't expect any request body.
   */
  modeOptionsGet(params?: {
  }): Observable<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: CmdOptions;
}> {

    return this.modeOptionsGet$Response(params).pipe(
      map((r: StrictHttpResponse<{
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: CmdOptions;
}>) => r.body as {
'success'?: boolean;
'error'?: string;
'raw'?: string;
'data'?: CmdOptions;
})
    );
  }

}
