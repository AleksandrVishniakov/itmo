package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.PhysicalMove;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.Type;

final public class SteelWing extends PhysicalMove {
  public SteelWing() {
    super(Type.STEEL, 70, 0.9);
  }

  @Override
  public void applySelfEffects(Pokemon p) {
    if (Math.random() <= 0.1) {
      System.out.println("Deffense increased!");
      p.addEffect(new Effect().stat(Stat.DEFENSE, 1));
    }
  }

  @Override
  public String describe() {
    return "uses ability \"Steel Wing\"";
  }
}
