package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.StatusMove;
import ru.ifmo.se.pokemon.Type;

final public class Growth extends StatusMove {
  public Growth() {
    super(Type.NORMAL, 0, 1);
  }

  @Override
  public void applySelfEffects(Pokemon p) {
    p.addEffect(new Effect().stat(Stat.ATTACK, 1).stat(Stat.SPECIAL_ATTACK, 1));
  }

  @Override
  public String describe() {
    return "uses ability \"Growth\"";
  }
}
