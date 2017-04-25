# lwdx
Lightweight document transformations. 
* LwDITA-based structured modular authoring in XML/XHTML/Markdown 
* Roundtrip transformations among them for authoring freedom 
* Tailored to handle XML mixed content 
* Avoids the horrors of Golang XML annotations 
## Contains 
* CLI commands
* API
## Features
* Relies on Golang encoder/xml package for parsing & tokenization
* Avoids as much XML & DITA complexity as possible
* Can make parse trees generated xmlx, etree, x2j, mxj, Go html
## Licensing
* LGPL 3.0 
* Dual licensing for unencumbered commercial use in the future (maybe)
## Lacks
* XML entity substitutions (but you were sposta use DITA conref's)
* Schema-driven WYSIWYG IBE (in-browser editor) (but this is roadmapped) 
