// https://docs.cypress.io/api/introduction/api.html

describe('My First Test', () => {
  it('check title and number of pages for first book', () => {
    cy.visit('/')
    cy.contains('tbody > tr:first-child > td:first-child', 'some-title')
    cy.contains('tbody > tr:first-child > td:nth-child(2)', '42')
  })

  it('check error state', () => {
    cy.visit('/', {
      onBeforeLoad (win) {
        // @ts-ignore
        win.__E2E_ERROR__ = true
      }
    })
    cy.contains('h3', 'Error')
  })

  it('check about page', () => {
    cy.visit('/')
    cy.get('a').contains('About').click()
    cy.contains('h1', 'This is an about page')
  })

})
