WALLET_ID=$(./stock_ballance -mode populate -entity wallet -description 'Taiar stocks' | tail -1)

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code mglu3 -name 'Magazine Luiza' | tail -1)
./stock_ballance -mode populate -entity asset -amount 212 -value 5423 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code suzb3 -name 'Suzano SA' | tail -1)
./stock_ballance -mode populate -entity asset -amount 205 -value 4470 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code itsa4 -name 'Itausa' | tail -1)
./stock_ballance -mode populate -entity asset -amount 383 -value 1353 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code klbn11 -name 'Klabin' | tail -1)
./stock_ballance -mode populate -entity asset -amount 200 -value 1541 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code abev3 -name 'Ambev SA' | tail -1)
./stock_ballance -mode populate -entity asset -amount 206 -value 1921 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code vvar3 -name 'Via Varejo' | tail -1)
./stock_ballance -mode populate -entity asset -amount 300 -value 1270 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code egie3 -name 'Engie' | tail -1)
./stock_ballance -mode populate -entity asset -amount 67 -value 5336 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code odpv3 -name 'Odontoprev' | tail -1)
./stock_ballance -mode populate -entity asset -amount 212 -value 1685 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code slce3 -name 'SLC Agrícola' | tail -1)
./stock_ballance -mode populate -entity asset -amount 143 -value 2467 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code flry3 -name 'Fleury' | tail -1)
./stock_ballance -mode populate -entity asset -amount 100 -value 3100 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code petr4 -name 'Petrobrás' | tail -1)
./stock_ballance -mode populate -entity asset -amount 100 -value 3033 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code b3sa3 -name 'B3' | tail -1)
./stock_ballance -mode populate -entity asset -amount 43 -value 4435 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code wege3 -name 'Wege' | tail -1)
./stock_ballance -mode populate -entity asset -amount 44 -value 3432 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code sqia3 -name 'Sínqia' | tail -1)
./stock_ballance -mode populate -entity asset -amount 60 -value 2450 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"

STOCK_ID=$(./stock_ballance -mode populate -entity stock -code mdia3 -name 'M. Dias Branco' | tail -1)
./stock_ballance -mode populate -entity asset -amount 27 -value 3969 -stock_id "$STOCK_ID" -wallet_id "$WALLET_ID"
