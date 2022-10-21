module.exports = function (config) {
    config.set({
        basePath: '',
        frameworks: ['jasmine', '@angular-devkit/build-angular'],
        plugins: [
            require('karma-jasmine'),
            require('karma-chrome-launcher'),
            require('karma-jasmine-html-reporter'),
            require('karma-coverage'),
            require('@angular-devkit/build-angular/plugins/karma'),
        ],
        client: {
            jasmine: {},
            clearContext: false,
        },
        jasmineHtmlReporter: {
            suppressAll: true,
        },
        coverageReporter: {
            dir: 'coverage',
            reporters: [
                { type: 'html', subdir: 'report-html' },
                { type: 'lcov', subdir: 'report-lcov' },
                { type: 'text-summary', subdir: 'report-text-summary' },
                { type: 'cobertura', subdir: 'report-cobertura' },
            ],
            fixWebpackSourcePaths: true,
        },
        reporters: ['progress', 'kjhtml', 'coverage'],
        port: 9876,
        colors: true,
        logLevel: config.LOG_INFO,
        autoWatch: true,
        browsers: ['Chrome', 'ChromeHeadless_NoSandbox'],
        singleRun: false,
        restartOnFileChange: true,
        customLaunchers: {
            ChromeHeadless_NoSandbox: {
                base: 'ChromeHeadless',
                flags: [
                    '--no-sandbox',
                    '--disable-setuid-sandbox',
                    '--lang=en_GB,en',
                    '--disable-gpu',
                    '--remote-debugging-port=9222',
                ],
            },
        },
    });
};
