# xml-injection (CWE-91)

La version vulnerable construit le document XML par concatenation de chaines avec `fmt.Sprintf` : une entree utilisateur contenant des balises XML modifie la structure du document et peut y ajouter des elements non prevus (ex: un role administrateur). La version corrigee utilise `encoding/xml.Marshal` sur une struct typee, qui echappe automatiquement les caracteres speciaux XML dans le contenu des elements.
