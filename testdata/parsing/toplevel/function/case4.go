package main
func f(prefix string, values ...int)
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(f)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('f')
    PsiElement(()('(')
    FunctionParameterListImpl
      FunctionParameterImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('prefix')
        PsiWhiteSpace(' ')
        TypeNameImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('string')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      FunctionParameterVariadicImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('values')
        PsiWhiteSpace(' ')
        PsiElement(...)('...')
        TypeNameImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('int')
    PsiElement())(')')
