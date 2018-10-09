// Copyright 2018 The Murilo Ijanc'. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cptec fornece métodos para obter históricos de medições meteorologia
// e dados das estações do CPTEC/INPE (Centro de Previsão de Tempo e Estudos
//Climáticos).
//
// Além da coleta desses dados o pacote fornece a exportação dos
// mesmos nos formatos csv e json.:
//
// Exemplo de saída csv:
//
//		id,uf,locality
//		32619,to,xambioa
//
// Exemplo de saída json:
//
//		{"id": "32619", "uf": "to", "locality": "xambioa"}
//
// O autor desenvolveu o pacote como forma de estudo e não se responsabiliza
// pelo uso dos dados.
//
// Aviso CPTEC/INPE:
//
// Os produtos apresentados nesta página não podem ser usados para propósitos
// comerciais, copiados integral ou parcialmente para a reprodução em meios de
// divulgação, sem a expressa autorização do CPTEC/INPE. Os usuários deverão
// sempre mencionar a fonte das informações e dados como "CPTEC/INPE" A geração
// e a divulgação de produtos operacionais obedecem critérios sistêmicos de
// controle de qualidade, padronização e periodicidade de disponibilização. Em
// nenhum caso o CPTEC/INPE pode ser responsabilizado por danos especiais,
// indiretos ou decorrentes, ou nenhum dano vinculado ao que provenha do uso
// destes produtos.
package cptec
