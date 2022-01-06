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
import { Powerstat } from '../models/powerstat';

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
  }): Observable<StrictHttpResponse<void>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.FrequencyPutPath, 'put');
    if (params) {
      rb.body(params.body, 'application/json');
    }

    return this.http.request(rb.build({
      responseType: 'text',
      accept: '*/*'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return (r as HttpResponse<any>).clone({ body: undefined }) as StrictHttpResponse<void>;
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
  }): Observable<void> {

    return this.frequencyPut$Response(params).pipe(
      map((r: StrictHttpResponse<void>) => r.body as void)
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
  }): Observable<StrictHttpResponse<void>> {

    const rb = new RequestBuilder(this.rootUrl, ApiService.PowerstatPutPath, 'put');
    if (params) {
      rb.body(params.body, 'application/json');
    }

    return this.http.request(rb.build({
      responseType: 'text',
      accept: '*/*'
    })).pipe(
      filter((r: any) => r instanceof HttpResponse),
      map((r: HttpResponse<any>) => {
        return (r as HttpResponse<any>).clone({ body: undefined }) as StrictHttpResponse<void>;
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
  }): Observable<void> {

    return this.powerstatPut$Response(params).pipe(
      map((r: StrictHttpResponse<void>) => r.body as void)
    );
  }

}
