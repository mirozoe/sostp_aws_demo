# Zdrojové kódy k tutoriálům
Následující materiály vznikly pro podporu vzdělávání v IT na Hotelové škole, Střední průmyslové škole a Obchodní akademii v Teplicích. Primárním zdrojem informací jsou video materiály a zde přístupné zdroje slouží jako podpora.

## O aplikaci
Zde je kompletní aplikace použitá ve video tutoriálu. Jedná se o velmi jednoduchou implementaci účetního deníku, která je hostována na platformě AWS. Front-end je JavaScript aplikace která komunikuje s back-end aplikací přes REST API.

## Obsah aresáře

### FE
V adresáří `fe` je zdrojový kód pro front-end aplikaci. Upload do AWS Amplify je prováděn automaticky po odeslání změn na GitHub.

### BE
Představuje BE aplikace, která je zodpovědná za příjem a odeslání dat HTTP klientovi a ukládání/čtení dat do AWS DynamoDB. Update aplikace je potřeba provést z adresáře be/ spuštěním skriptu `deploy.sh`
