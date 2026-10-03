package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.StatusMove;
import ru.ifmo.se.pokemon.Type;

final public class Acupressure extends StatusMove {
  public Acupressure() {
    super(Type.NORMAL, 0, 1);
  }

  @Override
  public void applySelfEffects(Pokemon p) {
    Effect allyBuff = new Effect();
    int chosenStat = (int) Math.round(Math.random() * (Stat.values().length - 1));

    allyBuff.stat(Stat.values()[chosenStat], 2);

    p.addEffect(allyBuff);
  }

  @Override
  public String describe() {
    return "uses ability \"Acupressure\"";
  }
}
