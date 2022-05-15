import { login } from '../utils.spec';

describe('Create baseString', () => {
    let email;
    let group;

    beforeEach(() => {
        email = Cypress.env('email');
        group = Cypress.env('group');
    });

    it('should create a baseString without translations', () => {
        login();
        addBasicInfo(group);
        cy.get('#baseStringSend').click();
        cy.get('#identifierBaseString').should('not.exist');
    });

    it('should create a baseString with translations', () => {
        login();
        addBasicInfo();
        addTranslation('es');
        addTranslation('en');
        cy.get('#baseStringSend').click();
        cy.get('#identifierBaseString').should('not.exist');
    });

    function addBasicInfo() {
        cy.get('#baseStringsMenuElement').click();
        cy.get('#createBaseStringsButton').click();
        const randomNumber = new Date().getTime();
        const identifier = `identifier-${randomNumber}`;
        Cypress.env('identifier', identifier);
        cy.get('#identifierBaseString').type(identifier).should('have.value', identifier);
        cy.get('#baseStringGroup').type(group).type('{enter}');
        cy.get('#baseStringLanguage').type('es').type('{enter}');
    }

    function addTranslation(isoCode) {
        cy.get('#openTranslationModalButton').click();
        const content = `content-${new Date().getTime()}`;
        const modalTranslationContent = cy.get('#modalTranslationContent');
        modalTranslationContent.type(content).should('have.value', content);
        cy.get('#modalTranslationStage').type('test').type('{enter}');
        cy.get('#modalTranslationLanguage').type(isoCode).type('{enter}');
        cy.get('#modalTranslationSend').click();
        modalTranslationContent.should('not.exist');
    }
});
