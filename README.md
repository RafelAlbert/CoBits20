# BitsXLaMarató20

# Inspiració

Atenent a la necessitat de garantir un entorn el més segur possible per a l'estudiantat, professorat i personal administratiu, en el marc de la pandèmia actual, hem decidit treballar en una aplicació per fer un seguiment de tots els possibles contagis de persones que hagin estat positives en una prova PCR.

# Què fa?

La nostra aplicació es fonamenta en dues vessants: una per l'estudiantat i l'altra pels gestors. 

**Per l'estudiantat**

En cada taula d'una aula hi haurà un codi QR, que l'estudiant escanejerà i, mitjançant l'API del Racó, l'estudiant s'identificarà amb les credencials pròpies del Racó mateix. Aleshores quedarà registrat per aquella taula concreta, d'aquella aula concreta, en aquella hora concreta.

**Pels gestors**

Quan es rebi la comunicació que una prova PCR realitzada a un estudiant o professor ha donat positiva, un dels gestors del sistema introduirà el nom d'usuari de la persona infectada i la data que es va realitzar la prova al sistema. Posteriorment, el sistema processarà tots els contactes sensibles susceptibles de ser infectats de la persona en qüestió, tots aquests contactes rebran una notificació per diversos canals, i hauran de confinar-se. Paral·lelament, els professors de les classes que tinguin com a mínim un alumne confinat també rebran un avís per saber que també han d'oferir l'opció de fer classes _online_.

A més, s'oferiran diverses opcions per consultar l'estat actual de qualsevol usuari.

# Com ho hem fet

Hem treballat per deduir el funcionament de la nova API del Racó per l'accés dels usuaris, 

# Reptes que ens hem trobat

El principal repte ha estat comprendre l'API del Racó de la FIB i posteriorment processar les dades obtingudes. Refrescar els nostres coneixements sobre interfícies i JavaScript també ha suposat una certa complicació.

# De què estem orgullosos

Aconseguir haver entès i comprès l'API del Racó per utilitzar-la en la identificació i obtenció de dades dels seus usuaris és motiu de satisfacció. Haver dissenyat una interfície senzilla però funcional, i ser capaços de processar i mostrar totes les dades amb les que treballem de forma adequada.

# Què hem après

Hem après a treballar amb l'API del Racó, tan en la seva versió antiga com amb la nova, i també hem treballat amb el protocol d'identificació OAuth. També hem expandit els nostres coneixement en disseny d'interfícies web i l'ús de JavaScript.

# Com seguir amb CoBits20

Es pot expandir el projecte de moltes maneres diferents, ja que sense cap mena de dubte podria esdevenir una eina molt útil per facilitar el funcionament el més normal possible de les classes a la universitat, mantenint uns criteris estrictes per garantir la seguretat de totes les persones involucrades.
