import { goLogin } from '../utils.spec';

describe('Login user', () => {
    let email;

    beforeEach(() => {
        email = Cypress.env('email');
    });

    it('should login with correct data', () => {
        const showLogin = goLogin();
        cy.get('#loginEmail').type(email).should('have.value', email);
        cy.get('#loginPassword').type(email).should('have.value', email);
        cy.get('#loginEnter').click();
        showLogin.should('not.exist');
    });

    it('should login with incorrect data', () => {
        const showLogin = goLogin();
        cy.get('#loginEmail').type(email).should('have.value', email);
        const loginEnter = cy.get('#loginEnter');
        loginEnter.should('be.disabled');
        cy.get('#loginPassword').type('incorrect').should('have.value', 'incorrect');
        loginEnter.click();
        showLogin.should('exist');
    });
});
