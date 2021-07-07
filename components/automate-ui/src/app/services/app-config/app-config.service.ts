import { HttpClient, HttpBackend } from '@angular/common/http';
import { Injectable } from '@angular/core';

interface BannerConfigTypes {
  show?: boolean;
  message?: string;
  background_color?: string;
  text_color?: string;
  idle_timeout?: number
}

interface ConfigTypes {
  banner?: BannerConfigTypes;
}

const initialConfig = {
<<<<<<< HEAD
    show: null,
    message: null,
    background_color: null,
    text_color: null,
    idle_timeout: null
=======
    banner: {
      show: null,
      message: null,
      background_color: null,
      text_color: null
    }
>>>>>>> d88dca54ce7f5744d47c68326ea4ccae31a4fa17
};

@Injectable({
  providedIn: 'root'
})

export class AppConfigService {

  public appConfig: ConfigTypes = initialConfig;

  constructor(private handler: HttpBackend) { }

  public loadAppConfig() {
    return new HttpClient(this.handler).get('/custom_settings.js')
      .toPromise()
      .then(data => this.appConfig = data)
      // .then(data => console.log(data, "dataaß"))
      // when there is no config, we can just reset the config to its initial empty values
      .catch(_error => this.appConfig = initialConfig);

  }

  get showBanner(): boolean {
    return this.appConfig.banner.show;
  }

  get bannerMessage(): string {
    return this.appConfig.banner.message;
  }

  get bannerBackgroundColor(): string {
    return this.convertToHex(this.appConfig.banner.background_color);
  }

  get idleTimeout(): number {
    return Number(this.appConfig.idle_timeout);
  }

  get bannerTextColor(): string {
    return this.convertToHex(this.appConfig.banner.text_color);
  }

  private convertToHex(color: string): string {
    return `#${color}`;
  }
}
