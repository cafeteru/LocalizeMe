// register.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test
describe('Register user', () => {
    it('should login with correct data', () => {
        const randomNumber = new Date().getTime();
        const email = `user${randomNumber}@email.es`;
        cy.visit('http://localhost:4200');
        const showLogin = cy.get('#showLogin');
        showLogin.click();
        cy.get('#loginRegister').click();
        cy.get('#registerEmail').type(email).should('have.value', email);
        cy.get('#registerPassword').type(email).should('have.value', email);
        cy.get('#registerCheckPassword').type(email).should('have.value', email);
        cy.get('#registerEnter').click();
        showLogin.should('not.exist');
    });
});
