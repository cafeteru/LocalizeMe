import { goLogin } from '../utils.spec';

describe('Register user', () => {
    it('should register with correct data', () => {
        const randomNumber = new Date().getTime();
        const email = `user${randomNumber}@email.es`;
        Cypress.env('email', email);
        const showLogin = goLogin();
        cy.get('#loginRegister').click();
        cy.get('#registerEmail').type(email).should('have.value', email);
        cy.get('#registerPassword').type(email).should('have.value', email);
        cy.get('#registerCheckPassword').type(email).should('have.value', email);
        cy.get('#registerEnter').click();
        showLogin.should('not.exist');
    });
});
